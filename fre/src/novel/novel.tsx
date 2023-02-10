import { h, useEffect, useState } from 'fre'
import { A, push } from '../use-route'
import { addChapter, getChapters, getNovel, getUser } from '../util/api'
import './novel.css'

export default function Novel(props) {
    const id = props.id

    const [novel, setNovel] = useState({} as any)
    const [data, setData] = useState({} as any)
    const [list, setList] = useState([])
    const [current, setIndex] = useState(0)

    useEffect(() => {
        getNovel(id).then(res => {
            setNovel(res.data)
            getChapters(res.data._id).then(res2 => {
                setList(res2.data)
            })
        })
    }, [])

    function changeData(key, val) {
        console.log(key, val)
        setData({
            ...data,
            [key]: val,
        })
    }


    function publish() {
        addChapter({
            ...data,
            status: '发布',
            nid: novel._id
        }).then(res => {
            alert(res.msg)
        })
    }

    function open(index) {
        setIndex(index)
    }

    const user = getUser()||{}

    const isAuthor = novel.aid === user._id

    return <div class='wrapper'>
        <div class='detail'>
            <h1>《{novel.title}》</h1>
            <p>{novel.content}</p>
            <ul class='info'>
                <li>{novel.status}</li>
                <li>{novel.size}</li>
                <li>{novel.aptitude}</li>
                <li>{novel.sort}</li>
            </ul>
            <ul class='tag'>
                {novel.tag && novel.tag.split(' ').filter(t => t.length > 0).map(tag => {
                    return <li>#{tag}</li>
                })}
                <li onClick={()=>push(`/publish/${novel._id}`)}>编辑小说</li>
            </ul>
        </div>

        <div class='list'>
            {list.map((item, index) => {
                const content = item.content.replace(/\s+/g,'\n')

                return <div class='chapter'>
                    <div onClick={() => {
                        open(index)
                    }}>
                        <h2>{item.title}</h2>
                        {index === current && <pre>{content}</pre>}
                    </div>
                </div>
            })}
        </div>

        <div class='reply'>
            {isAuthor && <input type="text" placeholder='请输入章节序号' onInput={e => changeData('oid', parseInt(e.target.value))} />}
            {isAuthor && <input type="text" placeholder='请输入标题' onInput={e => changeData('title', e.target.value)} />}
            <textarea name="" id="" rows="10" onInput={e => changeData('content', e.target.value.replace(/\s+/g,'\n'))}></textarea>
            <button onClick={publish}>发布</button>
        </div>
    </div>
}