import { h, useEffect, useState } from 'fre'

export default function List(props) {
    return <div class='panel'>
        <h1>{props.name}</h1>
        <ul>
            {props.list.map((item, index) => {
                // console.log(index, item)
                return <li class='no'>《{item.title}》<span>{item.bio}</span></li>
            })}
        </ul>
    </div>
}