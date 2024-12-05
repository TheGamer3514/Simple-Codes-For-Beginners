
using System;
using System.IO;

class Program
{
    static void Main()
    {
        string filePath = "example.txt";
        if (File.Exists(filePath))
        {
            string content = File.ReadAllText(filePath);
            Console.WriteLine(content);
        }
        else
        {
            Console.WriteLine("File not found.");
        }
    }
}
