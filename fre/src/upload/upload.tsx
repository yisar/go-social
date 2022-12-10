import { h, render } from 'fre'
import { getUser, loginPost } from '../util/api'
import { getAvatar } from '../util/avatar'
import './upload.css'

const tags = [["甜文","虐文","爽文",'狗血','意识流'],
['古代','现代','民国','未来'],
['HE','BE','OE'],
['1v1','NP','骨科','年上','年下','受转攻','直掰弯','攻控','受控'],
['快穿','悬疑','破镜重圆','强制爱','先虐受后虐攻','追妻'],
['ABO','生子','哨兵','支服'],
['娱乐圈','宫廷','网游'],
['霹雳','原神'],
['授权翻译','授权转载']]

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
        <ul class="tags">
            <li><h2>6.请选择标签</h2></li>
            {tags.map(group=><ul>
                {group.map(item=><li>{item}</li>)}
            </ul>)}
        </ul>
        <ul>
            <li><h2>7.请输入文案</h2></li>
            <li><input type="text" /></li>


        </ul>
    </div>
}
