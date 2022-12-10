import { get, post } from './post'

export function loginPost(user){
  return post(`https://www.htwxc.com/login`, user)
}