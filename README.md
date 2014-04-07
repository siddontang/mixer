# Mixer

Mixer is not only a mysql proxy, its aim to supply a simple solution for mysql use.

It aim to solve some mysql use problem in my company, which expects to support following feature:

- split read and write.
- mysql node HA, if main node crashed, switch to backup node automatically.

# Install 

    cd $WORKSPACE
    git clone git@github.com:siddontang/mixer.git src/github.com/siddontang/mixer
    
    cd src/github.com/siddontang/mixer

    ./bootstrap.sh

    . ./dev.env

    make
    make test



# Todo....

- supply custom rule to dispatch query to different mysql nodes. 
- statistics. 

# Feedback

now mixer is only simple and not perfect. I need your feedback to improve continually. Thank you very much!

Email: siddontang@gmail.com

QQ: 335498184

