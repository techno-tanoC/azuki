import style from './text_link.module.scss'

import Link from 'next/link'

export default function TextLink(props) {
  return (
    <Link href={props.href}>
      <a className={style.link}>
        {props.children}
      </a>
    </Link>
  )
}
