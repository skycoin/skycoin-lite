#Stop importing "github.com/skycoin/skycoin/src/util/logging", since it is incompatible with gopherjs.
grep -rl '\"github.com\/skycoin\/skycoin\/src\/util\/logging\"' ./vendor/github.com/skycoin/skycoin/src | xargs sed -i 's/\"github.com\/skycoin\/skycoin\/src\/util\/logging\"/ /g'
grep -rl 'var logger' ./vendor/github.com/skycoin/skycoin/src | xargs sed -i 's/var logger/\/\/var logger/g'
grep -rl 'logger =' ./vendor/github.com/skycoin/skycoin/src | xargs sed -i 's/logger =/\/\/logger =/g'

#Replace logger with log.
grep -rl 'logger\.' ./vendor/github.com/skycoin/skycoin/src | xargs sed -i 's/logger\./log\./g'

#Comment methods that existed in logger, but not in log.
grep -rl 'log\.Critical()' ./vendor/github.com/skycoin/skycoin/src | xargs sed -i 's/log\.Critical()/\/\/log\.Critical()/g'
grep -rl 'log\.Warning(' ./vendor/github.com/skycoin/skycoin/src | xargs sed -i 's/log\.Warning(/\/\/log\.Warning(/g'

#Import log where it is necessary.
goimports -w ./vendor/github.com/skycoin/skycoin/src