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

export function publishThread(data){
  return post(`https://www.cuipiya.net/thread/add`, data)
}

export function addPost(data){
  return post(`https://www.cuipiya.net/post/add`, data)
}

export function getThreads(sort){
  return get(`https://www.cuipiya.net/threads?sort=${sort}&page=1&pageSize=10`)
}

export function getThread(id){
  return get(`https://www.cuipiya.net/thread/detail/${id}`)
}

export function getPosts(nid){
  return get(`https://www.cuipiya.net/posts?nid=${nid}&page=1&pageSize=10`)
}