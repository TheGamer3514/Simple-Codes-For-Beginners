using Discord;
using Discord.WebSocket;
using System;
using System.Threading.Tasks;

namespace DiscordBot
{
    class Program
    {
        private static async Task Main(string[] args)
        {
            var token = Environment.GetEnvironmentVariable("TOKEN"); // Load the bot token from environment variables
            if (string.IsNullOrEmpty(token))
            {
                Console.WriteLine("Error: Bot token not found. Set the TOKEN environment variable.");
                return;
            }

            var client = new DiscordSocketClient();
            client.Log += Log; // Logging events
            client.MessageReceived += MessageReceived; // Handle messages

            await client.LoginAsync(TokenType.Bot, token);
            await client.StartAsync();

            // Block the program until it is closed
            await Task.Delay(-1);
        }

        private static Task Log(LogMessage log)
        {
            Console.WriteLine(log);
            return Task.CompletedTask;
        }

        private static async Task MessageReceived(SocketMessage arg)
        {
            if (arg.Author.IsBot) return; // Ignore bot messages

            if (arg.Content == "!ping")
            {
                await arg.Channel.SendMessageAsync("Pong!");
            }
        }
    }
}
