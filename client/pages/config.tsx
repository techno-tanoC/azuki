import ConfigLink from '../components/text_link'

import {useEffect, useState} from 'react'

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
      <ConfigLink href="/">
        top
      </ConfigLink>
    </div>
  )
}
