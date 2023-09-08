LINUX=false
OS="$(uname)"
if [[ "${OS}" == "Linux" ]]; then
  LINUX=true
fi

UNAME_MACHINE="$(/usr/bin/uname -m)"

echo $LINUX
echo $UNAME_MACHINE

if [[ $UNAME_MACHINE == "x86_64" ]] && [[ $LINUX == true ]]; then
  wget https://github.com/Tchoupinax/check-git/releases/download/v0.2.0/check-git_0.2.0_linux_amd64.tar.gz &> /dev/null
  tar xvf check-git_0.2.0_linux_amd64.tar.gz
  rm check-git_0.2.0_linux_amd64.tar.gz
fi
