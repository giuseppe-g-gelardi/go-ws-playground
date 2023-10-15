

import { URL } from '~/helpers/urlConstructor';
import { type User } from '~/types';


export async function GetAllUsers(): Promise<Array<User>> {
  return await fetch(URL('users')).then(res => res.json());
}

export async function GetUserById(id: string): Promise<User> {
  const res = await fetch(URL('user'), {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ _id: id })
  })
  return await res.json();
}



