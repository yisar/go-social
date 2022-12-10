import { get, post } from './post'

export function loginPost(user){
  return post(`https://www.htwxc.com/login`, user)
}

export function registerUser(user){
  return post(`https://www.htwxc.com/register`, user)
}

export function getUser(){
  return window.localStorage.getItem('author')
}