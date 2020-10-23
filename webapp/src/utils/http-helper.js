import config from '../config/config'

function ApiPost(url, body) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body)
    }

    let apiUrl = config.config.apiserver + url
    return fetch(apiUrl, requestOptions)
}

export default {
    ApiPost
}