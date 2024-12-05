// Import required modules
const { Client, GatewayIntentBits } = require('discord.js');
require('dotenv').config();

// Create a new client instance
const client = new Client({
    intents: [
        GatewayIntentBits.Guilds,
        GatewayIntentBits.GuildMessages,
        GatewayIntentBits.MessageContent
    ]
});

// Ready event
client.once('ready', () => {
    console.log(`Logged in as ${client.user.tag}!`);
});

// Message event listener
client.on('messageCreate', (message) => {
    // Ignore messages from bots
    if (message.author.bot) return;

    // Check for a command
    if (message.content === '!ping') {
        message.reply('Pong!');
    }
});

// Login to Discord
client.login(process.env.TOKEN).catch((error) => {
    console.error('Error logging in:', error);
});
