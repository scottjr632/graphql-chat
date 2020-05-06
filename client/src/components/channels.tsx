import React, { FC, useState, useEffect } from 'react'

import { useQuery } from '@apollo/react-hooks';
import { TabNav } from '@primer/components'

import CHANNELS_QUERY from '../queries/channelsQuery'
import { Channels, Channels_channels } from '../queries/types/Channels'

interface Props {
  onClick: (channelName: string) => any
}

const ChannelsTab: FC<Props> = ({ onClick }) => {
  const { data, loading } = useQuery<Channels>(CHANNELS_QUERY)

  const [activeChannel, setActiveChannel] = useState('');
  const [channels, setChannels] = useState<Channels_channels[]>([]);

  const handleClick = (name: string) => {
    onClick(name)
    setActiveChannel(name)
  }

  useEffect(() => {
    if (data?.channels && !loading) {
      setChannels(data.channels)
    }
  }, [data, loading])

  return (
    <TabNav aria-label="Main">
      {channels.map(channel => (
        <TabNav.Link onClick={() => handleClick(channel.name)} selected={activeChannel === channel.name}>
          {channel.name}
        </TabNav.Link>
      ))}
    </TabNav>
  )
}

export default ChannelsTab