># Notes: The tool sends passwords to OpenAI in plain text. Using this tool carries the risk of exposing your password. Please be aware of the potential risks.

# Password-Strength-Checker (psc)

Check the password strength via OpenAI

Follow up blog at [Revolutionizing Password Strength Evaluation with OpenAI: A New Open Source Project — PSC](https://towardsdev.com/revolutionizing-password-strength-evaluation-with-openai-a-new-open-source-project-psc-83a63af32bb6)

## Usage

### Prepare your OpenAI API Key

Way 1 is, set environment variable
```
export OPENAI_API_KEY="<replace with your openai api key>"
```

Way 2 is, put the key to local .env
```
mv .env.sample .env
```

update the API key with yours

### Run psc tool directly

Download the binary from [release](https://github.com/AbacusGPT/Password-Strength-Checker/releases)

Choice the suitable OS for your environment

uncompress and run it
```
# for example, from my macbook
$ wget https://github.com/AbacusGPT/Password-Strength-Checker/releases/download/v1.0.13/psc-1.0.13-darwin-amd64.tar.gz 

$ tar zxvf psc-1.0.13-darwin-amd64.tar.gz

$ mv psc-1.0.13-darwin-amd64 psc

$ ./psc --password 'Pass!234'
weak

$ ./psc --password '62sWJFk28gVnXK3u'
strong
```

Only for mac user, you need allow applications downloaded from "App Store and identified developers" in "Privacy & Security" setting.

![image](https://github.com/AbacusGPT/Password-Strength-Checker/assets/8954908/56f9c668-550b-47f0-9710-c2dde4e91444)


### build by yourself

* Clone this repo

```
git clone https://github.com/AbacusGPT/Password-Strength-Checker
```

* Prepare your OpenAI API Key

Way 1 is, set environment variable
```
export OPENAI_API_KEY="<replace with your openai api key>"
```

Way 2 is, put the key to local .env
```
mv .env.sample .env
```

update the API key with yours

* run the build

```
go build -o psc .
```

You should see a binary file `psc` generated locally

* run the test now
```
./psc --password 'Pass!234'
weak

$ ./psc --password '62sWJFk28gVnXK3u'
strong
```

### run tests
```
$ go test
Password 'Pass!234', Run 1: Result - strong
Password 'Pass!234', Run 2: Result - weak
Password 'Pass!234', Run 3: Result - weak
Password 'Pass!234', Run 4: Result - weak
Password 'Pass!234', Run 5: Result - weak
Password 'Pass!234', Run 6: Result - strong
Password 'Pass!234', Run 7: Result - strong
Password 'Pass!234', Run 8: Result - weak
Password 'Pass!234', Run 9: Result - weak
Password 'Pass!234', Run 10: Result - weak
Password '62sWJFk28gVnXK3u', Run 1: Result - strong
Password '62sWJFk28gVnXK3u', Run 2: Result - weak
Password '62sWJFk28gVnXK3u', Run 3: Result - strong
Password '62sWJFk28gVnXK3u', Run 4: Result - strong
Password '62sWJFk28gVnXK3u', Run 5: Result - strong
Password '62sWJFk28gVnXK3u', Run 6: Result - strong
Password '62sWJFk28gVnXK3u', Run 7: Result - strong
Password '62sWJFk28gVnXK3u', Run 8: Result - strong
Password '62sWJFk28gVnXK3u', Run 9: Result - strong
Password '62sWJFk28gVnXK3u', Run 10: Result - strong
PASS
ok  	github.com/AbacusGPT/Password-Strength-Checker	13.619s
```
