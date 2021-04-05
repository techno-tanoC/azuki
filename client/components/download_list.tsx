import Item from './download'
import style from './download_list.module.scss'

export default function DownloadList({ downloads, deleteItem }) {
  return (
    <div className={style.list}>
      {
        downloads.map(download => (
          <Item key={download.id} download={download} deleteItem={deleteItem} />
        ))
      }
    </div>
  )
}
