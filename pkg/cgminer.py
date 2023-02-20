import socket
import json
import sys


class CgminerAPI(object):
    """ Cgminer RPC API wrapper. """

    def __init__(self, host='localhost', port=4028):
        self.data = {}
        self.host = host
        self.port = port

    def command(self, command, arg=None):
        """ Initialize a socket connection,
        send a command (a json encoded dict) and
        receive the response (and decode it).
        """
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.settimeout(1)

        try:
            sock.connect((self.host, self.port))
            payload = {"command": command}
            if arg is not None:
                # Parameter must be converted to basestring (no int)
                payload.update({'parameter': arg})

            if sys.version_info.major == 2:
                sock.send(json.dumps(payload))
            if sys.version_info.major == 3:
                sock.send(bytes(json.dumps(payload), 'utf-8'))
            received = self._receive(sock)
        except Exception as e:
            return dict({'STATUS': [{'STATUS': 'error', 'description': "{}".format(e)}]})
        else:
            # the null byte makes json decoding unhappy
            # also add a comma on the output of the `stats` command by
            # replacing '}{' with '},{'
            try:
                return json.loads(received[:-1].replace('}{', '},{'))
            except Exception as e:
                return dict({'STATUS': [{'STATUS': 'error', 'description': "{}".format(e)}]})
        finally:
            # sock.shutdown(socket.SHUT_RDWR)
            sock.close()

    def _receive(self, sock, size=4096):
        msg = ''
        while 1:
            chunk = sock.recv(size)
            if chunk:
                if sys.version_info.major == 2:
                    msg += chunk
                if sys.version_info.major == 3:
                    msg += chunk.decode('utf-8')
            else:
                # end of message
                break
        return msg

    def __getattr__(self, attr):
        """ Allow us to make command calling methods.
        >>> cgminer = CgminerAPI()
        >>> cgminer.summary()
        """

        def out(arg=None):
            return self.command(attr, arg)

        return out


def get_summary(ip):
    cgminer = CgminerAPI(host=ip)
    output = cgminer.summary()
    output.update({"IP": ip})
    return dict(output)


def get_pools(ip):
    cgminer = CgminerAPI(host=ip)
    output = cgminer.pools()
    output.update({"IP": ip})
    return dict(output)


def get_stats(ip):
    cgminer = CgminerAPI(host=ip)
    output = cgminer.stats()
    output.update({"IP": ip})
    return dict(output)


if __name__ == '__main__':
    command = sys.argv[1]
    address = sys.argv[2]
    api = CgminerAPI(host=address)

    if command == 'summary':
        print(api.summary())
    elif command == 'stats':
        print(api.stats())