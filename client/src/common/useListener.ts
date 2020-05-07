import { useEffect, useState } from 'react';

type Formatter = (data: string) => any

let eventListener: EventSource

function useListener<T>(type: string, url: string, formatter?: Formatter) {
  const [event, setEvent] = useState<T | null>(null)

  useEffect(() => {
    if (!eventListener) {
      eventListener = new EventSource(url)
    }
    eventListener.addEventListener(type, (e: any) => {
      if (e) {
        const data = formatter ? formatter(e.data) : e.data
        setEvent(data as T)
      }
    })
  }, [formatter, type, url])

  return event
}

export default useListener