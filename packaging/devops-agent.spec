Name: devops-agent
Version: 1.0
Release: 1%{?dist}
Summary: DevOps Agent

License: MIT
BuildRequires: golang
Requires: python3

%install
install -D devops-agent %{buildroot}/usr/bin/devops-agent
install -D devops-agent.service %{buildroot}/usr/lib/systemd/system/devops-agent.service
