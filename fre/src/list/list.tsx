import { h, useEffect, useState } from 'fre'
import { A } from '../use-route'

export default function List(props) {
    return <div class='panel'>
        <h1>{props.name}</h1>
        <ul>
            {props.list.map((item, index) => {
                // console.log(index, item)
                return <A href={`/novel/ht${item._id}`}><li class='no'>《{item.title}》<span>{item.bio}</span></li></A>
            })}
        </ul>
    </div>
}