import {useEffect, useState} from 'react'
import Link from 'next/link'

export default function Config() {
  const [port, setPort] = useState("")

  useEffect(() => {
    const p = localStorage.getItem("port")
    setPort(p)
  }, [])

  const handleChange = event => {
    setPort(event.target.value)
  }

  const handleSubmit = (event) => {
    localStorage.setItem("port", port)
    event.preventDefault()
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          PORT:
          <input type="number" name="port" value={port} onChange={handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
      <Link href="/">
        <a>top</a>
      </Link>
    </div>
  )
}
