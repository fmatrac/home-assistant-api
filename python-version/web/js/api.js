// API Helper functions
const API = {
    baseUrl: '/api',

    async request(endpoint, options = {}) {
        const url = `${this.baseUrl}${endpoint}`;
        const config = {
            headers: {
                'Content-Type': 'application/json',
            },
            ...options,
        };

        if (config.body && typeof config.body === 'object') {
            config.body = JSON.stringify(config.body);
        }

        const response = await fetch(url, config);

        if (!response.ok) {
            const error = await response.json().catch(() => ({ error: 'Unknown error' }));
            throw new Error(error.error || 'Request failed');
        }

        if (response.status === 204) {
            return null;
        }

        return response.json();
    },

    // Wydarzenia
    wydarzenia: {
        getAll: () => API.request('/wydarzenia'),
        get: (id) => API.request(`/wydarzenia/get?id=${id}`),
        create: (data) => API.request('/wydarzenia', { method: 'POST', body: data }),
        update: (id, data) => API.request(`/wydarzenia?id=${id}`, { method: 'PUT', body: data }),
        delete: (id) => API.request(`/wydarzenia?id=${id}`, { method: 'DELETE' }),
    },

    // Przypomnienia
    przypomnienia: {
        getAll: (idWydarzenia = null) => {
            const params = idWydarzenia ? `?id_wydarzenia=${idWydarzenia}` : '';
            return API.request(`/przypomnienia${params}`);
        },
        get: (id) => API.request(`/przypomnienia/get?id=${id}`),
        create: (data) => API.request('/przypomnienia', { method: 'POST', body: data }),
        update: (id, data) => API.request(`/przypomnienia?id=${id}`, { method: 'PUT', body: data }),
        delete: (id) => API.request(`/przypomnienia?id=${id}`, { method: 'DELETE' }),
    },

    // Produkty
    produkty: {
        getAll: (kategoria = null) => {
            const params = kategoria ? `?kategoria=${encodeURIComponent(kategoria)}` : '';
            return API.request(`/produkty${params}`);
        },
        get: (id) => API.request(`/produkty/get?id=${id}`),
        getUlubione: () => API.request('/produkty/ulubione'),
        create: (data) => API.request('/produkty', { method: 'POST', body: data }),
        update: (id, data) => API.request(`/produkty?id=${id}`, { method: 'PUT', body: data }),
        delete: (id) => API.request(`/produkty?id=${id}`, { method: 'DELETE' }),
    },

    // Listy zakupów
    listyZakupow: {
        getAll: (status = null) => {
            const params = status ? `?status=${status}` : '';
            return API.request(`/listy-zakupow${params}`);
        },
        get: (id) => API.request(`/listy-zakupow/get?id=${id}`),
        create: (data) => API.request('/listy-zakupow', { method: 'POST', body: data }),
        update: (id, data) => API.request(`/listy-zakupow?id=${id}`, { method: 'PUT', body: data }),
        delete: (id) => API.request(`/listy-zakupow?id=${id}`, { method: 'DELETE' }),
    },

    // Pozycje listy zakupów
    pozycjeListy: {
        getByLista: (idListy) => API.request(`/pozycje-listy?id_listy=${idListy}`),
        get: (id) => API.request(`/pozycje-listy/get?id=${id}`),
        create: (data) => API.request('/pozycje-listy', { method: 'POST', body: data }),
        update: (id, data) => API.request(`/pozycje-listy?id=${id}`, { method: 'PUT', body: data }),
        delete: (id) => API.request(`/pozycje-listy?id=${id}`, { method: 'DELETE' }),
        markAsBought: (id) => API.request(`/pozycje-listy/kupione?id=${id}`, { method: 'POST' }),
    },

    // Stany magazynowe
    stanyMagazynowe: {
        getAll: (stan = null) => {
            const params = stan ? `?stan=${stan}` : '';
            return API.request(`/stany-magazynowe${params}`);
        },
        get: (id) => API.request(`/stany-magazynowe/get?id=${id}`),
        create: (data) => API.request('/stany-magazynowe', { method: 'POST', body: data }),
        update: (id, data) => API.request(`/stany-magazynowe?id=${id}`, { method: 'PUT', body: data }),
        delete: (id) => API.request(`/stany-magazynowe?id=${id}`, { method: 'DELETE' }),
    },

    // Historia stanu zapasów
    historiaZapasow: {
        getAll: (idProduktu = null) => {
            const params = idProduktu ? `?id_produktu=${idProduktu}` : '';
            return API.request(`/historia-zapasow${params}`);
        },
        create: (data) => API.request('/historia-zapasow', { method: 'POST', body: data }),
    },
};

