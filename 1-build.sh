#!/usr/bin/env bash

echo '
 __   __     __    __     ______    
/\ "-.\ \   /\ "-./  \   /\  ___\   
\ \ \-.  \  \ \ \-./\ \  \ \___  \  
 \ \_\\"\_\  \ \_\ \ \_\  \/\_____\ 
  \/_/ \/_/   \/_/  \/_/   \/_____/ 
                                    
'

echo "$(date) | Building Linux binary"
CGO_ENABLED=0 GOOS=linux go build -o ./bin/

echo "$(date) | Building Windows exe"
CGO_ENABLED=0 GOOS=windows go build -o ./bin/
