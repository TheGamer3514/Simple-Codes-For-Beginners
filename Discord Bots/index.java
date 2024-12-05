import net.dv8tion.jda.api.JDABuilder;
import net.dv8tion.jda.api.entities.Message;
import net.dv8tion.jda.api.events.message.MessageReceivedEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;

import javax.security.auth.login.LoginException;

public class DiscordBot extends ListenerAdapter {

    public static void main(String[] args) {
        String token = System.getenv("TOKEN"); // Load the bot token from environment variables
        if (token == null) {
            System.err.println("Error: Bot token not found. Set the TOKEN environment variable.");
            return;
        }

        try {
            JDABuilder.createDefault(token)
                    .addEventListeners(new DiscordBot())
                    .build();
        } catch (LoginException e) {
            System.err.println("Error logging in: " + e.getMessage());
        }
    }

    @Override
    public void onMessageReceived(MessageReceivedEvent event) {
        Message message = event.getMessage();
        String content = message.getContentRaw();

        // Ignore messages from bots
        if (event.getAuthor().isBot()) return;

        // Check for the !ping command
        if (content.equalsIgnoreCase("!ping")) {
            message.reply("Pong!").queue();
        }
    }
}
