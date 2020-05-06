import React, { FC } from 'react'
import { Timeline, StyledOcticon, Text, Flex, Heading } from '@primer/components'
import { Mail } from '@primer/octicons-react'

interface Message {
  content: string
  channelName: string
}

interface Props {
  messages: Message[]
}

const Messages: FC<Props> = ({ messages }) => {
  return (
    <Timeline style={{ flexDirection: 'column-reverse', overflow: 'hidden', overflowY: 'scroll' }}>

      {messages.map(message => (
        <Timeline.Item>
          <Timeline.Badge>
            <StyledOcticon icon={Mail} />
          </Timeline.Badge>
          <Timeline.Body>
              <Flex>
                <Heading fontSize={1} mb={2} color={'dark-grey'}>{message.channelName}</Heading>
                <Text as='p' paddingX={3} paddingY={1}>{message.content}</Text>
              </Flex>
          </Timeline.Body>
        </Timeline.Item>
      ))}

    </Timeline>
  )
}

export default Messages