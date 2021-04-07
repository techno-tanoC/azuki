import DownloadList from '../components/download_list'

import {useState, useEffect} from 'react'
import Link from 'next/link'

const fetchDownloads = async () => {
  const port = localStorage.getItem("port")
  const res = await fetch(`http://localhost:${port}/downloads`)
  const json = await res.json()
  return json
}

const deleteItem = (id: string) => {
  const port = localStorage.getItem("port")
  fetch(`http://localhost:${port}/downloads/${id}`, { method: "DELETE" })
}

export default function Index() {
  const [downloads, setDownloads] = useState([])

  useEffect(() => {
    const f = async () => {
      const news = await fetchDownloads()
      setDownloads(news)
    }

    f()
    const intervalId = setInterval(f, 1000)
    return () => clearInterval(intervalId)
  }, [])

  return (
    <div>
      <DownloadList downloads={downloads} deleteItem={deleteItem} />
      <Link href="/config">
        <a>config</a>
      </Link>
    </div>
  )
}
