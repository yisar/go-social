import { get, post } from './post'

export function loginPost(user){
  return post(`https://www.htwxc.com/login`, user)
}

export function registerUser(user){
  return post(`https://www.htwxc.com/register`, user)
}

export function getUser(){
  const author = window.localStorage.getItem('author')
  return JSON.parse(author)
}

export function publishNovel(data){
  return post(`https://www.htwxc.com/publish`, data)
}

export function getNoves(sort){
  return get(`https://www.htwxc.com/novels?sort=${sort}&page=1&pageSize=10`)
}