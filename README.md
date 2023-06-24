# Secure Touch (pingidentity.com) bot protection solutions

##  Introduction
The core functionality of this project is to generate valid SecuredTouchTokens, primarily for Asos. However, with minor tweaks, it should be able to support all SecuredTouch-protected websites. This tool has been used in production environments, handling over 1 million requests with an impressive 99.99% success rate.

## Acknowledgements
I am profoundly grateful to all the resources and individuals that contributed to my early journey into reverse engineering and coding. This project served as a stepping stone in refining my skills, and the lessons learned have been invaluable. It's important to note that as this was one of my early works, the Go code implemented in the project might not reflect the best practices.

Through the course of this project, I was able to identify and rectify my coding mistakes, particularly in writing Go code. The experience gleaned from working on this project has significantly enhanced my knowledge, leading to substantial improvements in my subsequent projects.

I am humbled to share this journey with you and my sincere thanks to everyone who has contributed to my development!

## What is SecuredTouch?
SecuredTouch is a leading solution in the field of behavioral biometrics for mobile transactions. It offers continuous authentication technologies that enhance security, reduce fraud, and improve the digital experience for users. By seamlessly collecting and analyzing a dynamic set of over 100 different behavioral parameters, such as keyboard typing and scrolling, SecuredTouch provides a robust and innovative security measure​1.

In this project, our primary focus is the generation of valid SecuredTouchTokens, which are integral to the security infrastructure that utilizes SecuredTouch's pioneering technologies.

## Features

#### Interaction

- **[Fake Mouse Generation](https://github.com/asavor/secure-touch/blob/main/generator/events/mact/mact.go)**
- **[Fake Keyboard Generation](https://github.com/asavor/secure-touch/blob/main/generator/interaction.go)**

#### Payload Encryption

- **[Payload Encryption](https://github.com/asavor/secure-touch/blob/main/generator/encryption.go)**
- **[Payload Encryption test](https://github.com/asavor/secure-touch/blob/main/generator/encrypt_test.go)**

#### Payload Type

- **[Starter](https://github.com/asavor/secure-touch/blob/main/generator/starter.go)**
- **[Metadata](https://github.com/asavor/secure-touch/blob/main/generator/metadata.go)**
- **[Interaction](https://github.com/asavor/secure-touch/blob/main/generator/interaction.go)**

#### Secure Touch Post

- **[Post Handler](https://github.com/asavor/secure-touch/blob/main/generator/secureTouch.go)**

#### Validate

- **[Generate valid SecuredTouchToken](https://github.com/asavor/secure-touch/blob/main/module/asos/login_test.go)**

## Contributions
We value and welcome contributions from the community. If you have any suggestions, issues, or improvements, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

