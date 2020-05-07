import React, { useEffect, useState, useMemo } from 'react';
import { useQuery, useSubscription } from '@apollo/react-hooks';

import MESSAGES from './queries/messagesQuery'
import MESSAGES_SUBSCRIPTION from './queries/messagesSubscription'
import { onMessageCreation } from './queries/types/onMessageCreation'
import { Messages, Messages_messages, MessagesVariables } from './queries/types/Messages'

import Channels from './components/channels'
import MessagesBox from './components/messages'
import InputBox from './components/inputBox'
import { Grid } from '@primer/components';
import useListener from './common/useListener';

const parseData = (msg: string) => JSON.parse(msg)

const GData = () => {
  const [activeChannel, setActiveChannel] = useState('');
  const [messages, setMessages] = useState<Messages_messages[]>([]);

  const messageSSE = useListener<Messages_messages>('message', 'http://localhost:8080/listen', parseData)

  const { loading, error, data, refetch } = useQuery<Messages, MessagesVariables>(MESSAGES, {
    variables: { channelName: activeChannel }, fetchPolicy: 'network-only'
  })
  const { data: createdMessage } = useSubscription<onMessageCreation>(MESSAGES_SUBSCRIPTION, {
    variables: { channelName: activeChannel }
  });

  useEffect(() => {
    if (messageSSE) {
      setMessages(prevMsgs => [...prevMsgs, { ...messageSSE, content: `${messageSSE.content} - From SSE`}])
    }
  }, [messageSSE])

  useEffect(() => {
    if (!loading && !error && data?.messages) {
      setMessages(data.messages)
    }
  }, [data, error, loading])

  useEffect(() => {
    if (createdMessage?.messageCreated) {
      setMessages(prevMessages => [...prevMessages, {...createdMessage.messageCreated, content: `${createdMessage.messageCreated.content} - From GraphQL websocket`}])
    }
  }, [createdMessage])

  const formattedMessages = useMemo(() => (
    messages.map(message => (
      {
        content: message.content,
        channelName: message.channel.name
      }
    )).reverse()
  ), [messages])

  return (
    <div className="App">
      <Grid gridTemplateRows={'3rem calc(100% - 6rem) 3rem'} height={'100vh'}>
        <Channels onClick={name => {
          setActiveChannel(name)
          refetch()
        }} />
        <MessagesBox messages={formattedMessages} />
        <InputBox onClick={() => {}} channelName={activeChannel} />
      </Grid>
    </div>
  )
}

export default GData