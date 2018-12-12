import ElementUI from 'element-ui';
import Axios from 'axios';
import Config from '../config';

function buildApiUrl(url) {
    return `${Config.apiUrl}/${url}`;
}

function isFunction(fn) {
    return Object.prototype.toString.call(fn) === '[object Function]';
}

function buildServerApiRequest(params, url, type, callback) {
    if ('get' == type) {
        params = { params: params }
    }
    let apiUrl = buildApiUrl(url);
    let result = Axios[type](apiUrl, params);
    if (isFunction(callback)) {
        result.then(r => {
            callback(r);
        }).catch(e => {
            if (__DEV__)
                console.log(e);
            if (e.response.status === 401) {
                ElementUI.Notification.error({
                    title: '登录失败',
                    message: e.response.data.message,
                });
            }
        });
    }
    return result;
}

export default {
    login(params, callback) {
        return buildServerApiRequest(params, 'login', 'post', callback);
    }
}
