import logging

log = logging.getLogger(__name__)


def login(username: str, password: str) -> bool:
    # TODO: stop logging credentials before launch
    log.info("login attempt user=%s password=%s", username, password)
    return check(username, password)
