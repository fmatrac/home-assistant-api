// App state
let currentSection = 'wydarzenia';
let produktyCache = [];
let wydarzeniaCache = [];
let currentCalendarDate = new Date();

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    initNavigation();
    loadProduktyCache();
    loadWydarzenia();
});

// Load produkty cache
async function loadProduktyCache() {
    try {
        produktyCache = await API.produkty.getAll() || [];
    } catch (err) {
        console.error('Failed to load produkty cache:', err);
    }
}

// Navigation
function initNavigation() {
    document.querySelectorAll('.nav-menu a').forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const section = link.dataset.section;
            switchSection(section);
        });
    });
}

function switchSection(section) {
    // Update nav
    document.querySelectorAll('.nav-menu a').forEach(a => a.classList.remove('active'));
    document.querySelector(`[data-section="${section}"]`).classList.add('active');

    // Update content
    document.querySelectorAll('.section').forEach(s => s.classList.remove('active'));
    document.getElementById(section).classList.add('active');

    currentSection = section;

    // Load data
    switch(section) {
        case 'wydarzenia': loadWydarzenia(); break;
        case 'przypomnienia': loadPrzypomnienia(); break;
        case 'produkty': loadProdukty(); break;
        case 'listy-zakupow': loadListyZakupow(); break;
        case 'stany-magazynowe': loadStanyMagazynowe(); break;
    }
}

// ==================== CALENDAR ====================
const MONTH_NAMES = [
    'Styczen', 'Luty', 'Marzec', 'Kwiecien', 'Maj', 'Czerwiec',
    'Lipiec', 'Sierpien', 'Wrzesien', 'Pazdziernik', 'Listopad', 'Grudzien'
];

function changeMonth(delta) {
    currentCalendarDate.setMonth(currentCalendarDate.getMonth() + delta);
    renderCalendar();
}

function renderCalendar() {
    const year = currentCalendarDate.getFullYear();
    const month = currentCalendarDate.getMonth();

    // Update header
    document.getElementById('calendar-month-year').textContent = `${MONTH_NAMES[month]} ${year}`;

    // Get first day of month and total days
    const firstDay = new Date(year, month, 1);
    const lastDay = new Date(year, month + 1, 0);
    const totalDays = lastDay.getDate();

    // Get day of week for first day (0 = Sunday, convert to Monday = 0)
    let startDayOfWeek = firstDay.getDay() - 1;
    if (startDayOfWeek < 0) startDayOfWeek = 6;

    // Get today for comparison
    const today = new Date();
    const todayStr = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}-${String(today.getDate()).padStart(2, '0')}`;

    // Build calendar days
    const container = document.getElementById('calendar-days');
    let html = '';

    // Previous month days
    const prevMonth = new Date(year, month, 0);
    const prevMonthDays = prevMonth.getDate();
    for (let i = startDayOfWeek - 1; i >= 0; i--) {
        const day = prevMonthDays - i;
        html += `<div class="calendar-day other-month"><span class="day-number">${day}</span></div>`;
    }

    // Current month days
    for (let day = 1; day <= totalDays; day++) {
        const dateStr = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
        const isToday = dateStr === todayStr;
        const dayEvents = getEventsForDate(dateStr);
        const hasEvents = dayEvents.length > 0;

        let classes = 'calendar-day';
        if (isToday) classes += ' today';
        if (hasEvents) classes += ' has-events';

        let eventDots = '';
        if (hasEvents) {
            eventDots = '<div class="event-dots">';
            dayEvents.slice(0, 3).forEach(e => {
                eventDots += `<div class="event-dot priority-${e.priorytet}"></div>`;
            });
            eventDots += '</div>';
        }

        html += `<div class="${classes}" onclick="showDayEvents('${dateStr}')">
            <span class="day-number">${day}</span>
            ${eventDots}
        </div>`;
    }

    // Next month days
    const totalCells = startDayOfWeek + totalDays;
    const remainingCells = totalCells % 7 === 0 ? 0 : 7 - (totalCells % 7);
    for (let day = 1; day <= remainingCells; day++) {
        html += `<div class="calendar-day other-month"><span class="day-number">${day}</span></div>`;
    }

    container.innerHTML = html;
}

function getEventsForDate(dateStr) {
    return wydarzeniaCache.filter(w => {
        const start = w.data_startu.split('T')[0];
        const end = w.data_zakonczenia.split('T')[0];
        return dateStr >= start && dateStr <= end;
    });
}

function showDayEvents(dateStr) {
    const events = getEventsForDate(dateStr);
    renderWydarzeniaList(events, dateStr);

    // Open modal to add event for this day
    showModalForDate('wydarzenie', dateStr);
}

// ==================== WYDARZENIA ====================
async function loadWydarzenia() {
    const container = document.getElementById('wydarzenia-list');
    container.innerHTML = '<div class="loading">Ladowanie...</div>';

    try {
        wydarzeniaCache = await API.wydarzenia.getAll() || [];
        renderCalendar();
        renderWydarzeniaList(wydarzeniaCache);
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

function renderWydarzeniaList(wydarzenia, selectedDate = null) {
    const container = document.getElementById('wydarzenia-list');

    if (!wydarzenia || wydarzenia.length === 0) {
        if (selectedDate) {
            container.innerHTML = `<div class="empty-state"><p>Brak wydarzen na ${selectedDate}</p></div>`;
        } else {
            container.innerHTML = '<div class="empty-state"><p>Brak wydarzen</p></div>';
        }
        return;
    }

    let header = selectedDate ? `<h4 style="margin-bottom:15px;">Wydarzenia na ${selectedDate}:</h4>` : '';

    container.innerHTML = header + wydarzenia.map(w => `
        <div class="item-card priority-${w.priorytet}">
            <h3>${escapeHtml(w.tytul)}</h3>
            <p>${escapeHtml(w.opis)}</p>
            <div class="meta">
                <span>${escapeHtml(w.miejsce)}</span>
                <span>${formatDate(w.data_startu)} - ${formatDate(w.data_zakonczenia)}</span>
                <span class="badge badge-${w.priorytet}">${w.priorytet}</span>
            </div>
            <div class="actions">
                <button class="btn btn-small btn-secondary" onclick="editWydarzenie(${w.id})">Edytuj</button>
                <button class="btn btn-small btn-danger" onclick="deleteWydarzenie(${w.id})">Usun</button>
            </div>
        </div>
    `).join('');
}

async function deleteWydarzenie(id) {
    if (!confirm('Czy na pewno chcesz usunac to wydarzenie?')) return;
    try {
        await API.wydarzenia.delete(id);
        loadWydarzenia();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function editWydarzenie(id) {
    try {
        const wydarzenie = await API.wydarzenia.get(id);
        showModal('wydarzenie', wydarzenie);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

// ==================== PRZYPOMNIENIA ====================
async function loadPrzypomnienia() {
    const container = document.getElementById('przypomnienia-list');
    container.innerHTML = '<div class="loading">Ladowanie...</div>';

    try {
        const przypomnienia = await API.przypomnienia.getAll();
        if (!przypomnienia || przypomnienia.length === 0) {
            container.innerHTML = '<div class="empty-state"><p>Brak przypomnien</p></div>';
            return;
        }

        container.innerHTML = przypomnienia.map(p => `
            <div class="item-card">
                <h3>${escapeHtml(p.tytul)}</h3>
                <p>${escapeHtml(p.opis)}</p>
                <div class="meta">
                    <span>${formatDate(p.nastepne_uruchomienie)}</span>
                    <span class="badge badge-${p.status}">${p.status}</span>
                </div>
                <div class="actions">
                    <button class="btn btn-small btn-secondary" onclick="editPrzypomnienie(${p.id})">Edytuj</button>
                    <button class="btn btn-small btn-danger" onclick="deletePrzypomnienie(${p.id})">Usun</button>
                </div>
            </div>
        `).join('');
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

async function deletePrzypomnienie(id) {
    if (!confirm('Czy na pewno chcesz usunac to przypomnienie?')) return;
    try {
        await API.przypomnienia.delete(id);
        loadPrzypomnienia();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function editPrzypomnienie(id) {
    try {
        const przypomnienie = await API.przypomnienia.get(id);
        showModal('przypomnienie', przypomnienie);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

// ==================== PRODUKTY ====================
let currentProduktyFilter = '';

async function loadProdukty() {
    const container = document.getElementById('produkty-list');
    container.innerHTML = '<div class="loading">Ladowanie...</div>';

    // Clear search
    const searchInput = document.getElementById('produkty-search');
    if (searchInput) searchInput.value = '';
    currentProduktyFilter = '';

    try {
        const produkty = await API.produkty.getAll();
        produktyCache = produkty || [];
        renderProdukty(produkty);
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

async function loadProduktyUlubione() {
    const container = document.getElementById('produkty-list');
    container.innerHTML = '<div class="loading">Ladowanie...</div>';

    // Clear search
    const searchInput = document.getElementById('produkty-search');
    if (searchInput) searchInput.value = '';
    currentProduktyFilter = '';

    try {
        const produkty = await API.produkty.getUlubione();
        renderProdukty(produkty);
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

function filterProdukty(query) {
    currentProduktyFilter = query.toLowerCase().trim();

    if (!currentProduktyFilter) {
        renderProdukty(produktyCache);
        return;
    }

    const filtered = produktyCache.filter(p =>
        p.nazwa.toLowerCase().includes(currentProduktyFilter) ||
        (p.kategoria && p.kategoria.toLowerCase().includes(currentProduktyFilter))
    );

    renderProdukty(filtered);
}

function renderProdukty(produkty) {
    const container = document.getElementById('produkty-list');

    if (!produkty || produkty.length === 0) {
        container.innerHTML = '<div class="empty-state"><p>Brak produktow</p></div>';
        return;
    }

    container.innerHTML = produkty.map(p => `
        <div class="item-card">
            <h3>
                <span class="favorite ${p.ulubiony ? 'active' : ''}" onclick="toggleUlubiony(${p.id}, ${!p.ulubiony})">
                    ${p.ulubiony ? '[*]' : '[ ]'}
                </span>
                ${escapeHtml(p.nazwa)}
            </h3>
            ${p.kategoria ? `<p>Kategoria: ${escapeHtml(p.kategoria)}</p>` : ''}
            ${p.link_do_kupna ? `<p><a href="${escapeHtml(p.link_do_kupna)}" target="_blank">Link do kupna</a></p>` : ''}
            <div class="actions">
                <button class="btn btn-small btn-secondary" onclick="editProdukt(${p.id})">Edytuj</button>
                <button class="btn btn-small btn-danger" onclick="deleteProdukt(${p.id})">Usun</button>
            </div>
        </div>
    `).join('');
}

async function toggleUlubiony(id, ulubiony) {
    try {
        const produkt = await API.produkty.get(id);
        produkt.ulubiony = ulubiony;
        await API.produkty.update(id, produkt);
        loadProdukty();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function deleteProdukt(id) {
    if (!confirm('Czy na pewno chcesz usunac ten produkt?')) return;
    try {
        await API.produkty.delete(id);
        loadProdukty();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function editProdukt(id) {
    try {
        const produkt = await API.produkty.get(id);
        showModal('produkt', produkt);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

// ==================== LISTY ZAKUPOW ====================
async function loadListyZakupow() {
    const container = document.getElementById('listy-zakupow-list');
    container.innerHTML = '<div class="loading">Ladowanie...</div>';

    try {
        const listy = await API.listyZakupow.getAll();
        if (!listy || listy.length === 0) {
            container.innerHTML = '<div class="empty-state"><p>Brak list zakupow</p></div>';
            return;
        }

        container.innerHTML = listy.map(l => `
            <div class="item-card lista-card" onclick="showPozycjeListy(${l.id}, '${escapeHtml(l.nazwa)}')" style="cursor:pointer;">
                <h3>${escapeHtml(l.nazwa)}</h3>
                <div class="meta">
                    <span class="badge badge-${l.status}">${l.status}</span>
                    <span>${formatDate(l.utworzono)}</span>
                </div>
                <div class="actions" onclick="event.stopPropagation()">
                    <button class="btn btn-small btn-secondary" onclick="editListaZakupow(${l.id})">Edytuj</button>
                    <button class="btn btn-small btn-danger" onclick="deleteListaZakupow(${l.id})">Usun</button>
                </div>
            </div>
        `).join('');
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

async function showPozycjeListy(idListy, nazwaListy) {
    const container = document.getElementById('listy-zakupow-list');
    container.innerHTML = '<div class="loading">Ladowanie pozycji...</div>';

    try {
        const pozycje = await API.pozycjeListy.getByLista(idListy);

        let html = `
            <div class="item-card">
                <h3>Lista: ${escapeHtml(nazwaListy)}</h3>
                <button class="btn btn-small btn-secondary" onclick="loadListyZakupow()">Powrot</button>
                <button class="btn btn-small btn-primary" onclick="showModal('pozycja', null, ${idListy}, '${escapeHtml(nazwaListy)}')">+ Dodaj pozycje</button>
                <hr style="margin: 15px 0;">
        `;

        if (!pozycje || pozycje.length === 0) {
            html += '<p>Brak pozycji na liscie</p>';
        } else {
            html += pozycje.map(p => `
                <div class="checkbox-item ${p.czy_kupione ? 'bought' : ''}">
                    <input type="checkbox" ${p.czy_kupione ? 'checked' : ''} 
                           onchange="toggleKupione(${p.id}, ${!p.czy_kupione}, ${idListy}, '${escapeHtml(nazwaListy)}')">
                    <span><strong>${getProduktNazwa(p.id_produktu)}</strong> x${p.ilosc}</span>
                    ${p.notatka ? `<span style="color:#95a5a6;"> - ${escapeHtml(p.notatka)}</span>` : ''}
                    <button class="btn btn-small btn-danger" style="margin-left:auto;" 
                            onclick="deletePozycja(${p.id}, ${idListy}, '${escapeHtml(nazwaListy)}')">x</button>
                </div>
            `).join('');
        }

        html += '</div>';
        container.innerHTML = html;
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

async function toggleKupione(id, kupione, idListy, nazwaListy) {
    try {
        if (kupione) {
            await API.pozycjeListy.markAsBought(id);
        } else {
            const pozycja = await API.pozycjeListy.get(id);
            pozycja.czy_kupione = false;
            await API.pozycjeListy.update(id, pozycja);
        }
        showPozycjeListy(idListy, nazwaListy);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function deletePozycja(id, idListy, nazwaListy) {
    if (!confirm('Usunac pozycje?')) return;
    try {
        await API.pozycjeListy.delete(id);
        showPozycjeListy(idListy, nazwaListy);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function deleteListaZakupow(id) {
    if (!confirm('Czy na pewno chcesz usunac te liste?')) return;
    try {
        await API.listyZakupow.delete(id);
        loadListyZakupow();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function editListaZakupow(id) {
    try {
        const lista = await API.listyZakupow.get(id);
        showModal('lista-zakupow', lista);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

function getProduktNazwa(id) {
    const produkt = produktyCache.find(p => p.id === id);
    return produkt ? produkt.nazwa : `Produkt #${id}`;
}

// ==================== STANY MAGAZYNOWE ====================
async function loadStanyMagazynowe(stan = null) {
    const container = document.getElementById('stany-magazynowe-list');
    container.innerHTML = '<div class="loading">Ladowanie...</div>';

    try {
        const stany = await API.stanyMagazynowe.getAll(stan);
        if (!stany || stany.length === 0) {
            container.innerHTML = '<div class="empty-state"><p>Brak stanow magazynowych</p></div>';
            return;
        }

        container.innerHTML = stany.map(s => `
            <div class="item-card">
                <h3>${getProduktNazwa(s.id_produktu)}</h3>
                <div class="meta">
                    <span class="badge badge-${s.stan}">${s.stan.toUpperCase()}</span>
                    <span>Sprawdzono: ${formatDate(s.ostatnio_sprawdzono)}</span>
                </div>
                <div class="actions">
                    <button class="btn btn-small btn-secondary" onclick="editStanMagazynowy(${s.id})">Edytuj</button>
                    <button class="btn btn-small btn-danger" onclick="deleteStanMagazynowy(${s.id})">Usun</button>
                </div>
            </div>
        `).join('');
    } catch (err) {
        container.innerHTML = `<div class="empty-state"><p>Blad: ${err.message}</p></div>`;
    }
}

async function deleteStanMagazynowy(id) {
    if (!confirm('Czy na pewno chcesz usunac ten stan?')) return;
    try {
        await API.stanyMagazynowe.delete(id);
        loadStanyMagazynowe();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function editStanMagazynowy(id) {
    try {
        const stan = await API.stanyMagazynowe.get(id);
        showModal('stan-magazynowy', stan);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

// ==================== AUTOCOMPLETE ====================
let autocompleteSelectedIndex = -1;

function setupAutocomplete(inputId, onSelect) {
    const input = document.getElementById(inputId);
    if (!input) return;

    // Create autocomplete container
    const container = input.parentElement;
    container.classList.add('autocomplete-container');

    // Create list
    let list = document.getElementById(inputId + '-autocomplete');
    if (!list) {
        list = document.createElement('div');
        list.id = inputId + '-autocomplete';
        list.className = 'autocomplete-list';
        list.style.display = 'none';
        container.appendChild(list);
    }

    input.addEventListener('input', () => {
        const value = input.value.trim().toLowerCase();
        if (value.length < 1) {
            list.style.display = 'none';
            return;
        }

        const matches = produktyCache.filter(p =>
            p.nazwa.toLowerCase().includes(value)
        ).slice(0, 10);

        if (matches.length === 0 && value.length > 0) {
            // Show "create new" option
            list.innerHTML = `
                <div class="autocomplete-item create-new" data-create="true" data-name="${escapeHtml(input.value)}">
                    + Dodaj nowy produkt: "${escapeHtml(input.value)}"
                </div>
            `;
            list.style.display = 'block';
        } else if (matches.length > 0) {
            let html = matches.map((p, i) => `
                <div class="autocomplete-item" data-id="${p.id}" data-name="${escapeHtml(p.nazwa)}">
                    ${escapeHtml(p.nazwa)}
                    ${p.kategoria ? `<span class="kategoria">${escapeHtml(p.kategoria)}</span>` : ''}
                </div>
            `).join('');

            // Add "create new" option if exact match not found
            const exactMatch = matches.find(p => p.nazwa.toLowerCase() === value);
            if (!exactMatch) {
                html += `
                    <div class="autocomplete-item create-new" data-create="true" data-name="${escapeHtml(input.value)}">
                        + Dodaj nowy produkt: "${escapeHtml(input.value)}"
                    </div>
                `;
            }

            list.innerHTML = html;
            list.style.display = 'block';
        } else {
            list.style.display = 'none';
        }

        autocompleteSelectedIndex = -1;
    });

    input.addEventListener('keydown', (e) => {
        const items = list.querySelectorAll('.autocomplete-item');
        if (items.length === 0) return;

        if (e.key === 'ArrowDown') {
            e.preventDefault();
            autocompleteSelectedIndex = Math.min(autocompleteSelectedIndex + 1, items.length - 1);
            updateAutocompleteSelection(items);
        } else if (e.key === 'ArrowUp') {
            e.preventDefault();
            autocompleteSelectedIndex = Math.max(autocompleteSelectedIndex - 1, 0);
            updateAutocompleteSelection(items);
        } else if (e.key === 'Enter' && autocompleteSelectedIndex >= 0) {
            e.preventDefault();
            items[autocompleteSelectedIndex].click();
        } else if (e.key === 'Escape') {
            list.style.display = 'none';
        }
    });

    list.addEventListener('click', async (e) => {
        const item = e.target.closest('.autocomplete-item');
        if (!item) return;

        if (item.dataset.create === 'true') {
            // Create new product
            const nazwa = item.dataset.name;
            try {
                const newProdukt = await API.produkty.create({ nazwa, ulubiony: false });
                produktyCache.push(newProdukt);
                onSelect(newProdukt.id, newProdukt.nazwa);
            } catch (err) {
                alert('Blad przy tworzeniu produktu: ' + err.message);
                return;
            }
        } else {
            onSelect(parseInt(item.dataset.id), item.dataset.name);
        }

        input.value = item.dataset.name;
        list.style.display = 'none';
    });

    // Hide on blur (with delay for click)
    input.addEventListener('blur', () => {
        setTimeout(() => {
            list.style.display = 'none';
        }, 200);
    });
}

function updateAutocompleteSelection(items) {
    items.forEach((item, i) => {
        item.classList.toggle('selected', i === autocompleteSelectedIndex);
    });
}

// ==================== MODAL ====================
let currentPozycjaListaId = null;
let currentPozycjaListaNazwa = null;
let selectedProduktId = null;
let prefilledDate = null;

function showModalForDate(type, dateStr) {
    prefilledDate = dateStr;
    showModal(type, null);
}

function showModal(type, data = null, extraParam = null, extraParam2 = null) {
    const modal = document.getElementById('modal');
    const title = document.getElementById('modal-title');
    const form = document.getElementById('modal-form');

    const isEdit = data !== null;
    selectedProduktId = null;

    // Prepare date values for wydarzenie
    let defaultStartDate = '';
    let defaultEndDate = '';
    if (prefilledDate && !isEdit) {
        defaultStartDate = `${prefilledDate}T09:00`;
        defaultEndDate = `${prefilledDate}T10:00`;
        prefilledDate = null; // Reset after use
    }

    switch(type) {
        case 'wydarzenie':
            title.textContent = isEdit ? 'Edytuj wydarzenie' : 'Nowe wydarzenie';
            form.innerHTML = `
                <div class="form-group">
                    <label>Tytul</label>
                    <input type="text" name="tytul" value="${data?.tytul || ''}" required>
                </div>
                <div class="form-group">
                    <label>Opis</label>
                    <textarea name="opis" required>${data?.opis || ''}</textarea>
                </div>
                <div class="form-group">
                    <label>Miejsce</label>
                    <input type="text" name="miejsce" value="${data?.miejsce || ''}" required>
                </div>
                <div class="form-group">
                    <label>Priorytet</label>
                    <select name="priorytet" required>
                        <option value="maly" ${data?.priorytet === 'maly' ? 'selected' : ''}>Maly</option>
                        <option value="sredni" ${data?.priorytet === 'sredni' ? 'selected' : ''}>Sredni</option>
                        <option value="duzy" ${data?.priorytet === 'duzy' ? 'selected' : ''}>Duzy</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Data rozpoczecia</label>
                    <input type="datetime-local" name="data_startu" value="${isEdit ? formatDateForInput(data?.data_startu) : defaultStartDate}" required>
                </div>
                <div class="form-group">
                    <label>Data zakonczenia</label>
                    <input type="datetime-local" name="data_zakonczenia" value="${isEdit ? formatDateForInput(data?.data_zakonczenia) : defaultEndDate}" required>
                </div>
                <div class="form-actions">
                    <button type="button" class="btn btn-secondary" onclick="closeModal()">Anuluj</button>
                    <button type="submit" class="btn btn-primary">${isEdit ? 'Zapisz' : 'Dodaj'}</button>
                </div>
            `;
            form.onsubmit = (e) => submitWydarzenie(e, data?.id);
            break;

        case 'przypomnienie':
            title.textContent = isEdit ? 'Edytuj przypomnienie' : 'Nowe przypomnienie';
            form.innerHTML = `
                <div class="form-group">
                    <label>Tytul</label>
                    <input type="text" name="tytul" value="${data?.tytul || ''}" required>
                </div>
                <div class="form-group">
                    <label>Opis</label>
                    <textarea name="opis" required>${data?.opis || ''}</textarea>
                </div>
                <div class="form-group">
                    <label>Status</label>
                    <select name="status" required>
                        <option value="aktywne" ${data?.status === 'aktywne' ? 'selected' : ''}>Aktywne</option>
                        <option value="zapauzowane" ${data?.status === 'zapauzowane' ? 'selected' : ''}>Zapauzowane</option>
                        <option value="zrealizowane" ${data?.status === 'zrealizowane' ? 'selected' : ''}>Zrealizowane</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Nastepne uruchomienie</label>
                    <input type="datetime-local" name="nastepne_uruchomienie" value="${formatDateForInput(data?.nastepne_uruchomienie)}" required>
                </div>
                <div class="form-actions">
                    <button type="button" class="btn btn-secondary" onclick="closeModal()">Anuluj</button>
                    <button type="submit" class="btn btn-primary">${isEdit ? 'Zapisz' : 'Dodaj'}</button>
                </div>
            `;
            form.onsubmit = (e) => submitPrzypomnienie(e, data?.id);
            break;

        case 'produkt':
            title.textContent = isEdit ? 'Edytuj produkt' : 'Nowy produkt';
            form.innerHTML = `
                <div class="form-group">
                    <label>Nazwa</label>
                    <input type="text" name="nazwa" value="${data?.nazwa || ''}" required>
                </div>
                <div class="form-group">
                    <label>Kategoria</label>
                    <input type="text" name="kategoria" value="${data?.kategoria || ''}">
                </div>
                <div class="form-group">
                    <label>Link do kupna</label>
                    <input type="url" name="link_do_kupna" value="${data?.link_do_kupna || ''}">
                </div>
                <div class="form-group">
                    <label>
                        <input type="checkbox" name="ulubiony" ${data?.ulubiony ? 'checked' : ''}> Ulubiony
                    </label>
                </div>
                <div class="form-actions">
                    <button type="button" class="btn btn-secondary" onclick="closeModal()">Anuluj</button>
                    <button type="submit" class="btn btn-primary">${isEdit ? 'Zapisz' : 'Dodaj'}</button>
                </div>
            `;
            form.onsubmit = (e) => submitProdukt(e, data?.id);
            break;

        case 'lista-zakupow':
            title.textContent = isEdit ? 'Edytuj liste' : 'Nowa lista zakupow';
            form.innerHTML = `
                <div class="form-group">
                    <label>Nazwa</label>
                    <input type="text" name="nazwa" value="${data?.nazwa || ''}" required>
                </div>
                <div class="form-group">
                    <label>Status</label>
                    <select name="status" required>
                        <option value="otwarta" ${data?.status === 'otwarta' ? 'selected' : ''}>Otwarta</option>
                        <option value="zamknieta" ${data?.status === 'zamknieta' ? 'selected' : ''}>Zamknieta</option>
                    </select>
                </div>
                <div class="form-actions">
                    <button type="button" class="btn btn-secondary" onclick="closeModal()">Anuluj</button>
                    <button type="submit" class="btn btn-primary">${isEdit ? 'Zapisz' : 'Dodaj'}</button>
                </div>
            `;
            form.onsubmit = (e) => submitListaZakupow(e, data?.id);
            break;

        case 'pozycja':
            currentPozycjaListaId = extraParam;
            currentPozycjaListaNazwa = extraParam2;
            title.textContent = 'Dodaj pozycje do listy';
            form.innerHTML = `
                <div class="form-group">
                    <label>Produkt (wpisz nazwe)</label>
                    <input type="text" id="produkt-search" placeholder="Szukaj lub dodaj nowy produkt..." autocomplete="off" required>
                    <input type="hidden" name="id_produktu" id="selected-produkt-id">
                </div>
                <div class="form-group">
                    <label>Ilosc</label>
                    <input type="number" name="ilosc" value="1" min="1" required>
                </div>
                <div class="form-group">
                    <label>Notatka</label>
                    <input type="text" name="notatka" placeholder="np. bez laktozy">
                </div>
                <div class="form-actions">
                    <button type="button" class="btn btn-secondary" onclick="closeModal()">Anuluj</button>
                    <button type="submit" class="btn btn-primary">Dodaj</button>
                </div>
            `;
            form.onsubmit = (e) => submitPozycja(e);

            // Setup autocomplete after form is in DOM
            setTimeout(() => {
                setupAutocomplete('produkt-search', (id, nazwa) => {
                    selectedProduktId = id;
                    document.getElementById('selected-produkt-id').value = id;
                });
            }, 0);
            break;

        case 'stan-magazynowy':
            title.textContent = isEdit ? 'Edytuj stan' : 'Nowy stan magazynowy';
            form.innerHTML = `
                <div class="form-group">
                    <label>Produkt</label>
                    <input type="text" id="stan-produkt-search" placeholder="Szukaj produkt..." ${isEdit ? 'disabled' : ''} value="${isEdit ? getProduktNazwa(data.id_produktu) : ''}" autocomplete="off" ${isEdit ? '' : 'required'}>
                    <input type="hidden" name="id_produktu" id="stan-selected-produkt-id" value="${data?.id_produktu || ''}">
                </div>
                <div class="form-group">
                    <label>Stan</label>
                    <select name="stan" required>
                        <option value="ok" ${data?.stan === 'ok' ? 'selected' : ''}>OK</option>
                        <option value="malo" ${data?.stan === 'malo' ? 'selected' : ''}>Malo</option>
                        <option value="brak" ${data?.stan === 'brak' ? 'selected' : ''}>Brak</option>
                    </select>
                </div>
                <div class="form-actions">
                    <button type="button" class="btn btn-secondary" onclick="closeModal()">Anuluj</button>
                    <button type="submit" class="btn btn-primary">${isEdit ? 'Zapisz' : 'Dodaj'}</button>
                </div>
            `;
            form.onsubmit = (e) => submitStanMagazynowy(e, data?.id, data?.id_produktu);

            if (!isEdit) {
                setTimeout(() => {
                    setupAutocomplete('stan-produkt-search', (id, nazwa) => {
                        document.getElementById('stan-selected-produkt-id').value = id;
                    });
                }, 0);
            }
            break;
    }

    modal.classList.add('active');
}

function closeModal() {
    document.getElementById('modal').classList.remove('active');
    selectedProduktId = null;
    currentPozycjaListaId = null;
    currentPozycjaListaNazwa = null;
    prefilledDate = null;
}

// ==================== FORM SUBMISSIONS ====================
async function submitWydarzenie(e, id = null) {
    e.preventDefault();
    const form = e.target;
    const data = {
        tytul: form.tytul.value,
        opis: form.opis.value,
        miejsce: form.miejsce.value,
        priorytet: form.priorytet.value,
        data_startu: new Date(form.data_startu.value).toISOString(),
        data_zakonczenia: new Date(form.data_zakonczenia.value).toISOString(),
    };

    try {
        if (id) {
            await API.wydarzenia.update(id, data);
        } else {
            await API.wydarzenia.create(data);
        }
        closeModal();
        loadWydarzenia();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function submitPrzypomnienie(e, id = null) {
    e.preventDefault();
    const form = e.target;
    const data = {
        tytul: form.tytul.value,
        opis: form.opis.value,
        status: form.status.value,
        nastepne_uruchomienie: new Date(form.nastepne_uruchomienie.value).toISOString(),
    };

    try {
        if (id) {
            await API.przypomnienia.update(id, data);
        } else {
            await API.przypomnienia.create(data);
        }
        closeModal();
        loadPrzypomnienia();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function submitProdukt(e, id = null) {
    e.preventDefault();
    const form = e.target;
    const data = {
        nazwa: form.nazwa.value,
        kategoria: form.kategoria.value || null,
        link_do_kupna: form.link_do_kupna.value || null,
        ulubiony: form.ulubiony.checked,
    };

    try {
        if (id) {
            await API.produkty.update(id, data);
        } else {
            await API.produkty.create(data);
        }
        closeModal();
        loadProdukty();
        loadProduktyCache();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function submitListaZakupow(e, id = null) {
    e.preventDefault();
    const form = e.target;
    const data = {
        nazwa: form.nazwa.value,
        status: form.status.value,
    };

    try {
        if (id) {
            await API.listyZakupow.update(id, data);
        } else {
            await API.listyZakupow.create(data);
        }
        closeModal();
        loadListyZakupow();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function submitPozycja(e) {
    e.preventDefault();
    const form = e.target;

    const idProduktu = parseInt(document.getElementById('selected-produkt-id').value);
    if (!idProduktu) {
        alert('Wybierz produkt z listy lub wpisz nazwe nowego produktu');
        return;
    }

    const notatkaValue = form.notatka.value.trim();

    const data = {
        id_listy_zakupow: currentPozycjaListaId,
        id_produktu: idProduktu,
        ilosc: parseInt(form.ilosc.value),
        notatka: notatkaValue || null,  // <- zmiana: jeśli puste to null
        czy_kupione: false,
    };

    try {
        await API.pozycjeListy.create(data);
        closeModal();
        showPozycjeListy(currentPozycjaListaId, currentPozycjaListaNazwa);
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

async function submitStanMagazynowy(e, id = null, existingIdProduktu = null) {
    e.preventDefault();
    const form = e.target;

    let idProduktu = existingIdProduktu;
    if (!idProduktu) {
        idProduktu = parseInt(document.getElementById('stan-selected-produkt-id').value);
        if (!idProduktu) {
            alert('Wybierz produkt z listy');
            return;
        }
    }

    const data = {
        id_produktu: idProduktu,
        stan: form.stan.value,
    };

    try {
        if (id) {
            await API.stanyMagazynowe.update(id, data);
        } else {
            await API.stanyMagazynowe.create(data);
        }
        closeModal();
        loadStanyMagazynowe();
    } catch (err) {
        alert('Blad: ' + err.message);
    }
}

// ==================== HELPERS ====================
function escapeHtml(text) {
    if (!text) return '';
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

function formatDate(dateStr) {
    if (!dateStr) return '';
    const date = new Date(dateStr);
    return date.toLocaleDateString('pl-PL', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

function formatDateForInput(dateStr) {
    if (!dateStr) return '';
    const date = new Date(dateStr);
    return date.toISOString().slice(0, 16);
}

