
using System;

class Program
{
    static void Main()
    {
        try
        {
            int number = int.Parse("NotANumber");
        }
        catch (FormatException e)
        {
            Console.WriteLine("Error: " + e.Message);
        }
    }
}
