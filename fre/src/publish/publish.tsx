import { h, render, useEffect, useState } from 'fre'
import { getNovel, getUser, loginPost, publishNovel } from '../util/api'
import { getAvatar } from '../util/avatar'
import './publish.css'

const tags = [["甜文", "虐文", "爽文", '狗血', '意识流'],
['古代', '现代', '民国', '未来'],
['HE', 'BE', 'OE'],
['1v1', 'NP', '骨科', '年上', '年下', '受转攻', '直掰弯', '攻控', '受控'],
['快穿', '悬疑', '破镜重圆', '强制爱', '先虐受后虐攻', '追妻'],
['ABO', '生子', '哨兵', '支服'],
['娱乐圈', '宫廷', '网游'],
['霹雳', '原神'],
['授权转载', '无版权转载']]

export default function Upload(props) {
    const [data, setData] = useState({
        tag: '', aid: getUser()._id,
    } as any)


    useEffect(() => {
        getNovel(props.id).then(res => {
            setData({
                ...data,
                aid: res.data.aid,
                ...res.data
            })
        })
    }, [])

    console.log(data.aid)

    function changeData(key, val) {
        setData({
            ...data,
            [key]: val,
        })
    }

    function selectTag(item) {
        if ((data.tag || '').indexOf(item) > -1) {
            setData({
                ...data,
                tag: data.tag.replace(` ${item}`, ''),
            })
        } else {
            setData({
                ...data,
                tag: data.tag + ' ' + item,
            })
        }

    }

    function publish() {
        if (Object.keys(data).length < 9) {
            alert('全部都要填！')
        }

        publishNovel(data).then(res => {
            alert(res.msg)
        })
    }
    return <div class="wrapper upload">
        <h1>Pubulsh Novel</h1>
        <ul>
            <li><h2>请输入题目</h2></li>
            <li><input type="text" onInput={e => changeData('title', e.target.value)} value={data.title} /></li>
        </ul>
        <ul>
            <li><h2>请输入封面</h2></li>
            <li><input type="text" onInput={e => changeData('thumb', e.target.value)} value={data.thumb} /></li>
        </ul>
        <ul>
            <li><h2>请选择分类</h2></li>
            <li><input type="radio" value="原创" name="sort" id='原创' onInput={e => changeData('sort', e.target.value)} checked={data.sort === "原创"} />
                <label htmlFor="原创">原创</label></li>
            <li><input type="radio" value="同人" name="sort" id='同人' onInput={e => changeData('sort', e.target.value)} checked={data.sort === "同人"} />
                <label htmlFor="同人">同人</label></li>
        </ul>
        <ul>
            <li><h2>请选择状态</h2></li>
            <li><input type="radio" name="status" value="完结" id='完结' onInput={e => changeData('status', e.target.value)} />
                <label htmlFor="完结">完结</label></li>
            <li><input type="radio" name="status" value="连载" id='连载' onInput={e => changeData('status', e.target.value)} />
                <label htmlFor="连载">连载</label></li>
            <li><input type="radio" name="status" value="暂停" id='暂停' onInput={e => changeData('status', e.target.value)} />
                <label htmlFor="暂停">暂停</label></li>
        </ul>
        <ul>
            <li><h2>请选择篇幅</h2></li>
            <li>            <input type="radio" value="短篇" name="size" id="短篇" onInput={e => changeData('size', e.target.value)} checked={data.size === "短篇"} />
                <label htmlFor="短篇">短篇</label></li>
            <li><input type="radio" value="长篇" name='size' id='长篇' onInput={e => changeData('size', e.target.value)} checked={data.size === "长篇"} />
                <label htmlFor="长篇">长篇</label></li>

        </ul>
        <ul>
            <li><h2>请选择性向</h2></li>
            <li><input type="radio" value="bl" name='aptitude' id='bl' onInput={e => changeData('aptitude', e.target.value)} />
                <label htmlFor="bl">bl</label></li>
            <li><input type="radio" value="bg" name='aptitude' id='bg' onInput={e => changeData('aptitude', e.target.value)} />
                <label htmlFor="bg">bg</label></li>
        </ul>
        <ul>
            <li><h2>请输入一句话简介</h2></li>
            <li><input type="text" onInput={e => changeData('bio', e.target.value)} value={data.bio} /></li>
        </ul>
        <ul class="tags">
            <li><h2>请选择标签</h2></li>
            {tags.map(group => <ul>
                {group.map((item, index) => <li onClick={() => selectTag(item)} key={index.toString()}
                    className={data.tag.indexOf(item) > -1 ? 'active' : ''}>{item}</li>)}
            </ul>)}
        </ul>
        <ul>
            <li><h2>请输入文案</h2></li>
            <li><input type="text" onInput={e => changeData('content', e.target.value)} value={data.content} /></li>
        </ul>

        <button onClick={publish}>发布</button>
    </div>
}
