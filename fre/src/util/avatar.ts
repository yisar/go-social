import md5 from 'blueimp-md5'

export function getAvatar(avatar) {
    let hash = md5(avatar)
    return `https://cdn.sep.cc/avatar/${hash}?s=100&d=retro`
}