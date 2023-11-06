import os
import types

from aiogram import *
from pytube import YouTube

bot = Bot("6517147428:AAG_OkAcpjQOolOjzXko3rjvE4lcwtFFWhU")
dp = Dispatcher(bot)

@dp.message_handler(commands=["start"])
async def start_message(message: types.Message):
    chat_id = message.chat.id
    await bot.send_message(chat_id,"Привет")

@dp.message_handler()
async def text_message(message:types.Message):
    chat_id = message.chat.id
    url = message.text
    yt = YouTube(url)
    if message.text.startswith == 'https://youtu.be/' or 'https://www.youtube.com/':
        await bot.send_message(chat_id, f"Downloading... {yt.title}"
                                        f"*Channel*:[{yt.author}]({yt.channel_url})",parse_mode="Markdown")
        await download_youtube_video(url,message,bot)

async def download_youtube_video(url, message, bot):
    yt = YouTube(url)
    stream = yt.streams.filter(progressive=True, file_extension="mp4")
    stream.get_highest_resolution().download(f'{message.chat.id}', f'{message.chat.id}_{yt.title}')
    with open(f"{message.chat.id}/{message.chat.id}_{yt.title}",'rb') as video:
        await bot.send_video(message.chat.id, video, caption="*Your video*",parse_mode="Markdown")

if __name__ == '__main__':
    executor.start_polling(dp)