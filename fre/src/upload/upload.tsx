import { h, render } from 'fre'
import { getUser } from '../util/api'
import { getAvatar } from '../util/avatar'
import './upload.css'

export default function Upload() {
    return <div class="wrapper upload">
        <h1>Pubulsh Novel</h1>
        <ul>
            <li><h2>1.请选择分类</h2></li>
            <li><input type="radio" name="type" value="yuanchuang" id="radio1" defaultChecked />
                <label htmlFor="radio1">原创</label></li>
            <li><input type="radio" name="type" value="tongren" id="radio2" />
                <label htmlFor="radio2">同人</label></li>
        </ul>
        <ul>
            <li><h2>2.请选择状态</h2></li>
            <li><input type="radio" name="type" value="yuanchuang" id="radio1" defaultChecked />
                <label htmlFor="radio1">完结</label></li>
            <li><input type="radio" name="type" value="tongren" id="radio2" />
                <label htmlFor="radio2">连载</label></li>
            <li><input type="radio" name="type" value="tongren" id="radio2" />
                <label htmlFor="radio2">暂停</label></li>
        </ul>
        <ul>
            <li><h2>3.请选择篇幅</h2></li>
            <li>            <input type="radio" name="type" value="yuanchuang" id="radio1" defaultChecked />
                <label htmlFor="radio1">短篇</label></li>
            <li><input type="radio" name="type" value="tongren" id="radio2" />
                <label htmlFor="radio2">长篇</label></li>

        </ul>
        <ul>
            <li><h2>4.请选择性向</h2></li>
            <li><input type="radio" name="type" value="yuanchuang" id="radio1" defaultChecked />
                <label htmlFor="radio1">男性向</label></li>
            <li><input type="radio" name="type" value="tongren" id="radio2" />
                <label htmlFor="radio2">女性向</label></li>
        </ul>
        <ul>
            <li><h2>5.请输入一句话简介</h2></li>
            <li><input type="text" /></li>


        </ul>
        <ul>
            <li><h2>6.请选择标签</h2></li>
        </ul>
        <ul>
            <li><h2>7.请输入文案</h2></li>
            <li><input type="text" /></li>


        </ul>
    </div>
}