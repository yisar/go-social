import { get, post } from './post'

export function loginPost(user){
  return post(`https://www.cuipiya.net/login`, user)
}

export function registerUser(user){
  return post(`https://www.cuipiya.net/register`, user)
}

export function getUser(){
  const user = window.localStorage.getItem('user')
  return JSON.parse(user)
}

export function getUserDetail(uid){
  return get(`https://www.cuipiya.net/user/detail/${uid}`)
}

export function publishThread(data){
  return post(`https://www.cuipiya.net/thread/add`, data)
}

export function addPost(data){
  return post(`https://www.cuipiya.net/post/add`, data)
}

export function updateUser(data){
  return post(`https://www.cuipiya.net/register`, data)
}

export function getThreads(sort){
  return get(`https://www.cuipiya.net/threads?sort=${sort}&page=1&pageSize=10`)
}

export function getThread(id){
  return get(`https://www.cuipiya.net/thread/detail/${id}`)
}

export function getPosts(tid){
  return get(`https://www.cuipiya.net/posts?tid=${tid}&page=1&pageSize=30`)
}

export function getPostDetail(id){
  return get(`https://www.cuipiya.net/post/detail/${id}`)
}
