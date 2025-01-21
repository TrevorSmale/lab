# Kickstart File for ts

# Basic system configuration
lang en_US.UTF-8
keyboard us
timezone America/New_York --utc
rootpw --iscrypted $6$examplehashedpassword
network --bootproto=dhcp --device=eth0
bootloader --location=mbr
clearpart --all --initlabel

# Partitioning
part /boot --fstype=xfs --size=512
part pv.01 --size=1 --grow
volgroup vg_root pv.01
logvol / --vgname=vg_root --size=10240 --name=root --fstype=xfs
logvol swap --vgname=vg_root --size=2048 --name=swap

# Packages
%packages
@core
openssh-clients   # Install SSH tools to generate keys
%end

# Post-installation script
%post
# Generate ED25519 key for the 'ts' user
mkdir -p /home/ts/.ssh
ssh-keygen -t ed25519 -f /home/ts/.ssh/id_ed25519 -q -N ""
chown -R ts:ts /home/ts/.ssh
chmod 700 /home/ts/.ssh
chmod 600 /home/ts/.ssh/id_ed25519
chmod 644 /home/ts/.ssh/id_ed25519.pub
%end