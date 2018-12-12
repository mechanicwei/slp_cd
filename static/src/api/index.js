import user from './user';
import deploy_server from './deploy_server';
import deploy_record from './deploy_record';

let allApi = Object.assign(
  user,
  deploy_server,
  deploy_record
);

export default allApi;
