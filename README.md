
PROBLEM STATEMENT

Now that you know more about the device we've created called an Orb, with the help of distributors, we are going to sign up and give away Worldcoin tokens to people around the world. For the purpose of this exercise, the signup works by submitting an iris image in a png format to the backend where it’s transformed into a hash and checked if it’s unique among the ones we've already seen.

We would like to ask you to write a "virtual orb" that could be used as a part of our testing infrastructure. You can imagine running it in a CI pipeline or being deployed as a workload in multiple regions and constantly validating our system.

The orb should:
- Periodically report status by calling an API and submitting battery, cpu usage, cpu temp, disk space
- Periodically simulate a signup and submit images to the API with an associated id

You can choose the language you feel most comfortable with, we have a slight preference for Java, Go, C++, Python. You have the freedom of deciding how the API will look like and we care about testing. This task will likely not take more than 4-5h and if it should, please try to prioritize the work and describe the work that you would accomplish if you would have more time. If you have any questions please feel free to reach out. 

Please submit your solution in the link below. 