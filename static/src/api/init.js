import ElementUI from 'element-ui';
import Axios from 'axios';
import Config from '../config';

function buildApiUrl(url) {
  return `${Config.apiUrl}/${Config.apiPrefix}/${url}`;
}

function setToken() {
  Axios.defaults.headers['Authorization'] = 'Bearer ' + sessionStorage.getItem(Config.tokenKey);
}

function isFunction(fn) {
  return Object.prototype.toString.call(fn) === '[object Function]';
}

function buildServerApiRequest(params, url, type, callback) {
  setToken();
  if ('get' == type) {
	  params={ params: params }
  }
  let apiUrl = buildApiUrl(url);
  let result = Axios[type](apiUrl, params);
  if (isFunction(callback)) {//没有回调则返回es6 promise
    result.then(r => {
      callback(r);
    }).catch(e => {
      if(__DEV__)
        console.log(e);
      ElementUI.Notification.error({
        title: '请求错误',
      });
    });
  }
  return result;
}

export function buildApiRequest(params, url, type, callback) {
  return buildServerApiRequest(params, url, type, callback);
}

