import DownloadList from '../components/download_list'

import {useState, useEffect} from 'react'

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

    const intervalId = setInterval(f, 1000)
    return () => clearInterval(intervalId)
  }, [])

  return (
    <DownloadList downloads={downloads} deleteItem={deleteItem} />
  )
}
