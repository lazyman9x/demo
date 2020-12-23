import axios from 'axios';
import {CurrentUser} from './CurrentUser';
import {SnackReporter} from './snack/SnackManager';

export const initAxios = (currentUser: CurrentUser, snack: SnackReporter) => {
    axios.interceptors.request.use((config) => {
        config.headers['X-Livechat-Key'] = currentUser.token();
        return config;
    });

    axios.interceptors.response.use(undefined, (error) => {
        if (!error.response) {
            snack('Livechat server is not reachable, try refreshing the page.');
            return Promise.reject(error);
        }

        const status = error.response.status;

        if (status === 401) {
            currentUser.tryAuthenticate().then(() => snack('Could not complete request.'));
        }

        if (status === 400 || status === 403 || status === 500) {
            snack(error.response.data.error + ': ' + error.response.data.errorDescription);
        }

        return Promise.reject(error);
    });
};
