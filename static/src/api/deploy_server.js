import { buildApiRequest as Request } from './init'

export default {
    fetchAll: (data, c) => {
        Request({}, `deploy_repos/${data.repo_id}/servers`, 'get', c)
    },
    update: (data, c) => {
        Request(data, `deploy_servers/${data.id}`, 'patch', c)
    },
    create: (data, c) => {
        Request(data, `deploy_servers`, 'post', c)
    }
}
