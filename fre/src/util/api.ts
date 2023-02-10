import { get, post } from './post'

export function loginPost(user){
  return post(`https://www.cuipiya.net/login`, user)
}

export function registerUser(user){
  return post(`https://www.cuipiya.net/register`, user)
}

export function getUser(){
  const author = window.localStorage.getItem('author')
  return JSON.parse(author)
}

export function publishNovel(data){
  return post(`https://www.cuipiya.net/novel/add`, data)
}

export function addChapter(data){
  return post(`https://www.cuipiya.net/chapter/add`, data)
}

export function getNovels(sort){
  return get(`https://www.cuipiya.net/novels?sort=${sort}&page=1&pageSize=10`)
}

export function getNovel(id){
  return get(`https://www.cuipiya.net/novel/detail/${id}`)
}

export function getChapters(nid){
  return get(`https://www.cuipiya.net/chapters?nid=${nid}&page=1&pageSize=10`)
}