import { h, useEffect, useState } from 'fre'
import { getNoves } from '../util/api'

export default function Home() {
    const [list, setList] = useState([])
    useEffect(() => {
        getNoves('原创').then((res:any) => {
            setList(res.data)
        })
    }, [])
    console.log(list)
    return <div>
        {list.map((item, index) => {
            // console.log(index, item)
            return <div>{item.title}</div>
        })}
    </div>
}