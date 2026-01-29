// API Helper functions
const API = {
    baseUrl: '/api',

    getToken() {
        return localStorage.getItem('token');
    },

    setToken(token) {
        localStorage.setItem('token', token);
    },

    removeToken() {
        localStorage.removeItem('token');
    },

    isAuthenticated() {
        return !!this.getToken();
    },

    getTokenPayload() {
        const token = this.getToken();
        if (!token) return null;
        try {
            const base64Url = token.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const jsonPayload = decodeURIComponent(atob(base64).split('').map(c =>
                '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
            ).join(''));
            return JSON.parse(jsonPayload);
        } catch (e) {
            return null;
        }
    },

    isAdmin() {
        const payload = this.getTokenPayload();
        return payload && payload.rola === 'admin';
    },

    logout() {
        this.removeToken();
        window.location.href = '/login';
    },

    async request(endpoint, options = {}) {
        const url = `${this.baseUrl}${endpoint}`;
        const token = this.getToken();

        const config = {
            headers: {
                'Content-Type': 'application/json',
            },
            ...options,
        };

        // Add Authorization header if token exists
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }

        if (config.body && typeof config.body === 'object') {
            config.body = JSON.stringify(config.body);
        }

        const response = await fetch(url, config);

        // Handle 401 Unauthorized - redirect to login
        if (response.status === 401) {
            this.removeToken();
            window.location.href = '/login';
            return;
        }

        if (!response.ok) {
            const error = await response.json().catch(() => ({ error: 'Unknown error' }));
            throw new Error(error.detail || error.error || 'Request failed');
        }

        if (response.status === 204) {
            return null;
        }

        return response.json();
    },

    // Auth
    auth: {
        login: async (email, password) => {
            const response = await fetch(`${API.baseUrl}/auth/login/json`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, haslo: password }),
            });
            const data = await response.json();
            if (!response.ok) throw new Error(data.detail || 'Login failed');
            API.setToken(data.access_token);
            return data;
        },
        register: (data) => API.request('/auth/register', { method: 'POST', body: data }),
        me: () => API.request('/auth/me'),
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

    // Admin
    admin: {
        getUsers: () => API.request('/admin/users'),
        getUser: (id) => API.request(`/admin/users/${id}`),
        updateUser: (id, data) => API.request(`/admin/users/${id}`, { method: 'PUT', body: data }),
        deleteUser: (id) => API.request(`/admin/users/${id}`, { method: 'DELETE' }),
        getStats: () => API.request('/admin/stats'),
    },
};

