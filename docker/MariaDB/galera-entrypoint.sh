#!/usr/bin/env bash

[ "$DEBUG" == 'true' ] && set -x

set -eo pipefail

CONTAINER_ALREADY_STARTED="CONTAINER_ALREADY_STARTED"
if [ ! -e $CONTAINER_ALREADY_STARTED ]; then
    echo "First container startup"
    touch $CONTAINER_ALREADY_STARTED
    if [ -n "$WSREP_NEW_CLUSTER" ]; then
        echo "New cluster init"
        set -- "$@" --wsrep-new-cluster
    else 
        # Workaround to skip mysql db initialization.
        # It looks like the only way is to create an empty /var/lib/mysql/mysql directory! 
        # That's completely crazy. it took me hours to figure it out...
        echo 'Skip MySql initialization'
        mkdir -p /var/lib/mysql/mysql
    fi
else
    echo "Not first container startup"
fi

if [ -n "$WSREP_CLUSTER_ADDRESS" -a "$1" == 'mysqld' ]; then

	echo '>> Creating Galera Config'
	export MYSQL_INITDB_SKIP_TZINFO="yes"
	export MYSQL_ALLOW_EMPTY_PASSWORD="yes"

	cat <<- EOF > /etc/mysql/conf.d/galera-auto-generated.cnf
	# Galera Cluster Auto Generated Config
	[server]
	bind-address="0.0.0.0"
	binlog_format="row"
	default_storage_engine="InnoDB"
	innodb_autoinc_lock_mode="2"
	innodb_locks_unsafe_for_binlog="1"

	[galera]
	wsrep_on="on"
	wsrep_provider="${WSREP_PROVIDER:-/usr/lib/libgalera_smm.so}"
	wsrep_provider_options="${WSREP_PROVIDER_OPTIONS}"
	wsrep_cluster_address="${WSREP_CLUSTER_ADDRESS}"
	wsrep_cluster_name="${WSREP_CLUSTER_NAME:-my_wsrep_cluster}"
	wsrep_node_name="${WSREP_NODE_NAME:-$(hostname -s)}"
	wsrep_sst_auth="${WSREP_SST_AUTH}"
	wsrep_sst_method="${WSREP_SST_METHOD:-rsync}"
	EOF

	if [ -n "$WSREP_NODE_ADDRESS" ]; then
    		echo wsrep_node_address="${WSREP_NODE_ADDRESS}" >> /etc/mysql/conf.d/galera-auto-generated.cnf
	fi

elif [ -n "$WSREP_CLUSTER_ADDRESS" -a "$1" == 'garbd' ]; then

	echo '>> Configuring Garbd'
	set -- "$@" --address=$WSREP_CLUSTER_ADDRESS --group=${WSREP_CLUSTER_NAME:-my_wsrep_cluster} --name=$(hostname)

fi

exec /usr/local/bin/docker-entrypoint.sh "$@"
 
