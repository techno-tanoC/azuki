import DownloadList from '../components/download_list'

import {useState, useEffect} from 'react'

const fetchDownloads = async (endpoint: string) => {
  const url = new URL("/downloads", endpoint).toString()
  const res = await fetch(url)
  const json = await res.json()
  return json
}

const deleteItem = (endpoint: string) => {
  return (id: string) => {
    const url = new URL(`downloads/${id}`, endpoint).toString()
    fetch(url, { method: "DELETE" })
  }
}

export default function Index({ endpoint }) {
  const [downloads, setDownloads] = useState([])

  useEffect(() => {
    const f = async () => {
      const news = await fetchDownloads(endpoint)
      setDownloads(news)
    }

    f()
    const intervalId = setInterval(f, 1000)
    return () => clearInterval(intervalId)
  }, [])

  return (
    <div>
      <DownloadList downloads={downloads} deleteItem={deleteItem(endpoint)} />
    </div>
  )
}

export function getServerSideProps() {
  return {
    props: {
      endpoint: process.env.API_ENDPOINT ?? "http://localhost:8080"
    }
  }
}
