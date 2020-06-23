#!/usr/bin/python3
import discord
import getpass

class MyClient(discord.Client):
  async def on_ready(self):
    print('Logged on as {0}!'.format(self.user))

  async def on_message(self, message):
     print('Message from {0.author}: {0.content}'.format(message))

token = getpass.getpass(prompt='What is your Bot Token? ')
client = MyClient()
client.run(token)
