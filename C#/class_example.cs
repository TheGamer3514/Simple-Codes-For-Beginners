
using System;

class Car
{
    public string Brand;
    public int Speed;

    public void Drive()
    {
        Console.WriteLine(Brand + " is driving at " + Speed + " km/h.");
    }
}

class Program
{
    static void Main()
    {
        Car car = new Car();
        car.Brand = "Toyota";
        car.Speed = 100;
        car.Drive();
    }
}
