import React, { FC, useState, useCallback } from 'react'

import { TextInput, Grid, ButtonPrimary } from '@primer/components'
import { useMutation } from '@apollo/react-hooks';

import CREATE_MESSAGE from '../queries/createMessage'
import { CreateMessage, CreateMessageVariables } from '../queries/types/CreateMessage';

interface Props {
  channelName: string
  onClick: (message: string) => any
}

const InputBox: FC<Props> = ({ onClick, channelName }) => {
  const [input, setInput] = useState('');
  const [createMessage] = useMutation<CreateMessage, CreateMessageVariables>(CREATE_MESSAGE, {
    variables: { channelName, content: input }
  })

  const handleClick = useCallback(() => {
    onClick(input)
    createMessage()
    setInput('')
  }, [createMessage, input, onClick])

  return (
    <Grid gridTemplateColumns="80% 20%" gridGap={1} padding={'0 1rem'} marginBottom={'0.5rem'}>
      <TextInput value={input} aria-label="Message" name="message" placeholder="Message" onChange={e => setInput(e.currentTarget.value)} />
      <ButtonPrimary onClick={handleClick}>Send</ButtonPrimary>
    </Grid>
  )
}

export default InputBox