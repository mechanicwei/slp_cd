import { buildApiRequest as Request } from './init'

export default {
    fetchAll: (params, c) => {
        let server_id = params.server_id
        delete  params.server_id
        Request(params, `deploy_servers/${server_id}/records`, 'get', c)
    }
}
