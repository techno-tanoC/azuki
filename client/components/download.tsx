import style from './download.module.scss'

export default function Download({
  download, deleteItem
}) {
  const {id, name, total, size} = download
  const percent = `${Math.floor(size * 100 / total)}%`
  const ratio = `${size.toLocaleString()} / ${total.toLocaleString()}`

  return (
    <div className={style.card}>
      <div className={style.body}>
        <span style={{ width: percent }} className={style.progress} />
        <div className={style.name}>
          {name}
        </div>
        <div className={style.counts}>
          {percent}
          <br />
          {ratio}
        </div>
      </div>
      <button className={style.button} onClick={() => deleteItem(id)}>
        cancel
      </button>
    </div>
  )
}
