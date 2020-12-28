FROM python:3.6-stretch

COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt

ENV TZ=US/Arizona
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY main.py /tmp/main.py

CMD ["python", "-u", "/tmp/main.py"]