import HttpHelper from '../utils/http-helper'

export default class ProjectService {

    async AddProject(project) {
        let res = await HttpHelper.ApiPost("/api/project", project)
        if (res.status !== 200) {
            throw Error("Unable to save")
        }
    }
}