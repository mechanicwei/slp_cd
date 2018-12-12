import { buildApiRequest as Request } from './init'

export default {
    fetchAll: (c) => Request({}, 'deploy_repos', 'get', c),
    update: (data, c) => {
        Request(data, `deploy_repos/${data.id}`, 'patch', c)
    },
    create: (data, c) => {
        data.id = undefined
        Request(data, `deploy_repos`, 'post', c)
    }
}
