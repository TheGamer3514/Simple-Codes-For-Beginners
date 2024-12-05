import discord
from discord.ext import commands
import os

# Load bot token from environment variable
TOKEN = os.getenv('TOKEN')

# Create the bot instance with a command prefix
bot = commands.Bot(command_prefix="!", intents=discord.Intents.default())

# Event: Bot is ready
@bot.event
async def on_ready():
    print(f"Logged in as {bot.user}!")

# Command: !ping
@bot.command()
async def ping(ctx):
    await ctx.send("Pong!")

# Run the bot
if TOKEN:
    bot.run(TOKEN)
else:
    print("Error: Bot token not found. Make sure to set the TOKEN environment variable.")
